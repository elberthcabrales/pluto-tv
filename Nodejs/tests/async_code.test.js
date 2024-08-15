const axios = require('axios');
const makeRequests = require('../Concurrent_requests/concurrent_requests');
jest.mock('axios');


describe('makeRequests', () => {
    it('Change this code so instead of doing 10 requests one after the other, execute them all at once and log the response:', async () => {
        const mockData = {
            data: 'mock data',
        };

        axios.get.mockResolvedValue(mockData);

        await makeRequests();

        expect(axios.get).toHaveBeenCalledTimes(10);
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data1');
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data2');
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data3');
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data4');
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data5');
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data6');
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data7');
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data8');
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data9');
        expect(axios.get).toHaveBeenCalledWith('https://api.example.com/data10');
    });
})
