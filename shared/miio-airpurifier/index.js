const miio = require('miio');
miio.device({ address: '', token: "" })
    .then(async (device) => {
        const res = await device.call("get_prop", ['temp_dec', "aqi", "humidity"])
        process.stdout.write(JSON.stringify({
            aqi: res[1],
            temp_dec: res[0],
            humidity: res[2],
        }))
        process.exit(0)
    })
    .catch(err => {
        console.error(err)
    });