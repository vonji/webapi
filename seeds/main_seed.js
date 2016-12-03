const moment = require('moment');

exports.seed = function (knex, Promise) {
    return knex('users').del()
    .then(function () {
        return Promise.all([
            knex('users').insert({
                id: 1,
                email: 'loup.peluso@vonji.fr',
                password: 'loup.peluso',
                displayed_name: 'Loup',
                real_name: 'Loup Peluso',
                description: 'Coucou les p\'tits lapins.',
                motto: "J'aime les pommes",
                birthday: moment('1985-04-03'),
                created_at: moment(),
                updated_at: moment()
            }),
        ]);
    });
};