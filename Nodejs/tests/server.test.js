const request = require('supertest');
const app = require('../Express_api/server');

const moviesDb = [
    { id: 1, title: "Inception", director: "Christopher Nolan" },
    { id: 2, title: "The Matrix", director: "Lana Wachowski, Lilly Wachowski" },
    { id: 3, title: "Interstellar", director: "Christopher Nolan" },
];

describe('GET /movies', () => {
    it('should return a list of movies', async () => {
        const res = await request(app).get('/movies');
        expect(res.status).toBe(200);
        expect(Array.isArray(res.body)).toBe(true);
        expect(res.body).toHaveLength(3);
        expect(res.body[0]).toEqual(moviesDb[0]);
    });
});

describe('POST /movies', () => {
    it('should add a new movie to the list', async () => {
        const newMovie = { title: "The Dark Knight", director: "Christopher Nolan" };
        const res = await request(app).post('/movies').send(newMovie);
        expect(res.status).toBe(201);
        expect(res.body).toEqual(expect.objectContaining({ id: 4, ...newMovie }));

        const getRes = await request(app).get('/movies');
        expect(getRes.body).toHaveLength(4);
    });
});

describe('GET /movies/:id', () => {
    it('should return the details of a movie by ID', async () => {
        const res = await request(app).get('/movies/1');
        expect(res.status).toBe(200);
        expect(res.body).toEqual(moviesDb[0]);
    });

    it('should return 404 if the movie is not found', async () => {
        const res = await request(app).get('/movies/999');
        expect(res.status).toBe(404);
        expect(res.body).toEqual({ message: 'Movie not found' });
    });
});
