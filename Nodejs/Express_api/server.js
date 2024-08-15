const express = require('express');
const app = express();
app.use(express.json());

const moviesDb = [
    { id: 1, title: "Inception", director: "Christopher Nolan" },
    { id: 2, title: "The Matrix", director: "Lana Wachowski, Lilly Wachowski" },
    { id: 3, title: "Interstellar", director: "Christopher Nolan" },
];

// GET /movies - Returns a list of movies
app.get('/movies', (req, res) => {
    res.json(moviesDb);
});

// POST /movies - Adds a new movie to the list
app.post('/movies', (req, res) => {
    const newMovie = {
        id: moviesDb.length + 1,
        title: req.body.title,
        director: req.body.director
    };
    moviesDb.push(newMovie);
    res.status(201).json(newMovie);
});

// GET /movies/:id - Returns the details of a movie by ID
app.get('/movies/:id', (req, res) => {
    const movieId = parseInt(req.params.id, 10);
    const movie = moviesDb.find(m => m.id === movieId);

    if (!movie) {
        return res.status(404).json({ message: 'Movie not found' });
    }

    res.json(movie);
});

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server is running on port ${PORT}`);
});

module.exports = app;
