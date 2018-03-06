var fs = require('fs');

var data = JSON.parse(fs.readFileSync('from.json'))


let countTypes = ['DNS_TIMERTT_FROM_3WHS(S->D)','RTT_FULL_SZ_MIN(S->D)','RTT_FULL_SZ_MAX(S->D)']

data = data.PORTS;

let myData = {};



for (let countType of countTypes){
    myData[countType] = [];
    for(let item of data){
        if(item[countType]){
            item[countType]["PORT"] = item.PORT
            myData[countType].push(item[countType])
        }
    }
}
console.log(myData);









function groupBy(data,func){
    let groups = {};
    data.forEach(i=>{
        let group = func(i);
        groups[group] = groups[group] || [];
        groups[group].push(i);
    });
    return groups;
}


// export function groupBy<T>(data:Array<T>,func:(item:T)=>any):Map<any,Array<T>>{
//     let groups:Map<any,Array<T>> = new Map<any,Array<T>>();
//     data.forEach(i=>{
//         let group = func(i);
//         groups[group] = groups[group] || new Array<T>()
//         groups[group].push(i)
//     })
//     return groups;
// }