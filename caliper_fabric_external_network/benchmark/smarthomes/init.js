/*
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

'use strict';

module.exports.info  = 'Creating Devices.';

let txIndex = 0;
let temperatures = ['30', '32', '12', '35', '36', '40', '125'];
let owners = ['Salman', 'Sindre', 'Dennis', 'Tony Stark'];
let bc, contx;

module.exports.init = function(blockchain, context, args) {
    bc = blockchain;
    contx = context;

    return Promise.resolve();
};

module.exports.run = function() {
    txIndex++;
    let deviceId = 'DEVICE_' + txIndex.toString() + '_' + process.pid.toString();
    let deviceTemp = temperatures[txIndex % temperatures.length];
    let deviceHome = 'HOME_'+ txIndex.toString() + '_' + process.pid.toString(); 
    let deviceOwner = owners[txIndex % owners.length];

    let args;
    if (bc.bcType === 'fabric-ccp') {
        args = {
            chaincodeFunction: 'initDevice',
            chaincodeArguments: [deviceId, deviceTemp, deviceHome, deviceOwner],
        };
    } else {
        args = {
            verb: 'initDevice',
            id: deviceId,
            deviceReading: deviceTemp,
            home: deviceHome,
            owner: deviceOwner
        };
    }

    return bc.invokeSmartContract(contx, 'device', 'v11.2', args, 30);
};

module.exports.end = function() {
    return Promise.resolve();
};