//	Copyright 2024 Mike Nguyen <hey(at)mike.ee>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// JS shim to launch the relevant binary.

const os = require('os');
const child_process = require('child_process');

function getPlatformAndArch() {
  let platform = os.platform();
  let arch = os.arch();

  // Rename win32 to windows
  if (platform === 'win32') {
    platform = 'windows';
  }

  // Rename x64 to amd64
  if (arch === 'x64') {
    arch = 'amd64';
  }

  return { platform, arch };
}

function main() {
  const { platform, arch } = getPlatformAndArch();

  // Build the executable path based on platform and architecture
  let fullPath = './out/main';
  fullPath += `-${platform}-${arch}`;

  // Use child_process to execute the command
  child_process.exec(fullPath, (error, stdout, stderr) => {
    if (error) {
      console.error(`Error executing command: ${error}`);
      return;
    }
    console.log(`stdout: ${stdout}`);
    console.error(`stderr: ${stderr}`);
  });
}

if (require.main === module) {
    main()
}
