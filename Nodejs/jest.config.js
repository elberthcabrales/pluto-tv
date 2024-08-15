module.exports = {
    testEnvironment: 'node',
    testMatch: ['**/tests/**/*.js?(x)', '**/?(*.)+(spec|test).js?(x)'],
    moduleFileExtensions: ['js', 'json'],
    transform: {
        '^.+\\.js?$': 'babel-jest',
    },
};