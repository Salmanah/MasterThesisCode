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

module.exports.info  = 'Sending device reading.';


let bc, contx;


module.exports.init = function(blockchain, context, args) {
    bc = blockchain;
    contx = context;
    return Promise.resolve();
};

module.exports.run = function() {

    let args;
    if (bc.bcType === 'fabric-ccp') {
        args = {
            chaincodeFunction: 'sendDeviceReading'
        };
    }else{
        args = {
            verb: 'sendDeviceReading',
            DeviceID: "DEVICE_001",
            Type: "Fridge",
            Data: `xEy7Nh0OJBIGH34nQZ8fRd/U9mBF79mh1D4VXrEO8g4K1JzYvq+/QEPIMkv0Cn8QUpX0jBvnVHoz
            okKntYBmr3L5+VPqpvdrPQwWEZ9TNrmJH7Vkgbp6Wof4Iq94QM0PsRzwIrrXAv5NDUYoaVbCFvBX
            p2vjd+Dhw4U74bYcKU9kVhCejEpiw325CCjeZX/L50kk49JW/KL/rYWWBrGCfQt7akLIP3onKRPJ
            779Qy1zGrGRHDnaiM66BHivjSsJafjf69+BKgtL2aq1QHzoZVaK48y5HVPC6Qp2H+mPQ3qxr5QG2
            K7aH9eRYDvGk1uXtIs4JOBmisOCv1UbttgmYKxg3qp86fkplFn7mKvLw7k/ZlmEkhdWgzUXVBmHn
            5ZzFQRxtgdLMOpq0rK3xCTqkBuEf7kQ9DGUhHjF0XPs1j3TgXufutx+1rcWkTpO6Vs9l2qU5BfZA
            Mnq8pzO+5t/I1RlJJ/GekmWm4MxgxjE4qOYVDwN5pFs0oJuqCMnI4pypYuIBNYEz0I5fm7huvrWt`

        };
    }
    return bc.invokeSmartContract(contx, 'device', 'v0', args, 30);
};

module.exports.end = function() {
    return Promise.resolve();
};
