const path = require('path');
const fs = require('fs');

import New from '../src/cli/new'

describe('New command', () =>
{
    test('can create a new request file', () => {
        const new = new New();

		try
		{
			new.action('GET', 'posts/1');
		} catch(e: Exception) {
			expect(e).toMatch('error');
		}
    });
});
