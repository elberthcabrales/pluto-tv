const axios = require('axios');
const fetchData = require('../Asynchronous_code/async_code');
jest.mock('axios');

describe('fetchData', () => {
    it('Rewrite the function below from promises to use async/await style. Use a fake API endpoint', async () => {
        const mockData = {
            userId: 1,
            id: 1,
            title: 'delectus aut autem',
            completed: false,
        };

        axios.get.mockResolvedValue({ data: mockData });

        const data = await fetchData();

        expect(data).toEqual(mockData);
        expect(axios.get).toHaveBeenCalledWith('https://jsonplaceholder.typicode.com/todos/1');
    });
});
