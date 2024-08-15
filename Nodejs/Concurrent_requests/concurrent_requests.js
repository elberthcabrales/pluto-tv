const axios = require('axios');

async function makeRequests() {
    const urls = [
        'https://api.example.com/data1',
        'https://api.example.com/data2',
        'https://api.example.com/data3',
        'https://api.example.com/data4',
        'https://api.example.com/data5',
        'https://api.example.com/data6',
        'https://api.example.com/data7',
        'https://api.example.com/data8',
        'https://api.example.com/data9',
        'https://api.example.com/data10',
    ];

    const requests = urls.map(url => axios.get(url));

    try {
        const responses = await Promise.all(requests);

        responses.forEach(response => {
            console.log(response.data);
        });
    } catch (error) {
        console.error('Error during requests:', error);
    }
}

module.exports = makeRequests;