const localDnsMap = {
  'localhost': '127.0.0.1'
}

const localDns = (host: string) => {
  var re = new RegExp(Object.keys(localDnsMap).join("|"), "gi");

  return host.replace(re, (matched) => {
    return localDnsMap[matched.toLowerCase()];
  });
}

export default localDns;