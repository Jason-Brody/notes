let regexValue = '\/\\d{8}\/.+'

let regex =  new RegExp(regexValue);
const str = `bytedance-oss-ies-abroad-debug-data-singapore/20180210/HYPSTAR-ID-7f51a130-f952-45c5-8801-639605a18297.HAR`;
//console.log(regex.test(str));
//regex.compile();
console.log(regex.exec(str).entries().next().value[1]);



