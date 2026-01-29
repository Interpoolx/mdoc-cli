const { execSync } = require('child_process');
const fs = require('fs');
const path = require('path');

const version = process.argv[2];

if (!version) {
    console.error('Usage: node scripts/publish-version.js <version>');
    process.exit(1);
}

console.log(`Starting publication for version: ${version}`);

try {
    // 1. Update package.json version
    const pkgPath = path.join(__dirname, '..', 'package.json');
    const pkg = JSON.parse(fs.readFileSync(pkgPath, 'utf8'));
    pkg.version = version;
    fs.writeFileSync(pkgPath, JSON.stringify(pkg, null, 2) + '\n');
    console.log('✓ Updated package.json version');

    // 2. Compile Go binaries
    console.log('Compiling Go binaries...');
    const platforms = [
        { os: 'windows', arch: 'amd64', output: 'bin/mdoc-cli-windows-amd64.exe' },
        { os: 'linux', arch: 'amd64', output: 'bin/mdoc-cli-linux-amd64' },
        { os: 'darwin', arch: 'arm64', output: 'bin/mdoc-cli-darwin-arm64' }
    ];

    platforms.forEach(p => {
        console.log(`  - Building for ${p.os}/${p.arch}...`);
        let cmd = `go build -o ${p.output}`;
        if (p.os !== process.platform) {
            if (process.platform === 'win32') {
                cmd = `$env:GOOS='${p.os}'; $env:GOARCH='${p.arch}'; go build -o ${p.output}`;
            } else {
                cmd = `GOOS=${p.os} GOARCH=${p.arch} go build -o ${p.output}`;
            }
        }

        // Use powershell on windows, else use default shell
        const shell = process.platform === 'win32' ? 'powershell.exe' : undefined;
        let finalCmd = cmd;
        if (process.platform === 'win32') {
            // Refresh path and run command
            finalCmd = `$env:Path = [System.Environment]::GetEnvironmentVariable("Path","Machine") + ";" + [System.Environment]::GetEnvironmentVariable("Path","User"); ${cmd}`;
        }
        execSync(finalCmd, { stdio: 'inherit', shell });
    });
    console.log('✓ Compiled Go binaries');

    // 3. Publish to NPM
    console.log('Publishing to NPM...');
    execSync('npm publish --access public', { stdio: 'inherit' });
    console.log('✓ Published to NPM');

    console.log(`\nSuccessfully published version ${version}!`);

} catch (error) {
    console.error('\n✗ Publication failed:');
    console.error(error.message);
    process.exit(1);
}
