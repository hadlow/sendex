const path = require('path');
const fs = require('fs');

import Init from '../src/cli/init'

describe('Init command', () =>
{
    test('can create the config file', () => {
        const init = new Init();

		try
		{
			init.action();
		} catch(e: Exception) {
			expect(e).toMatch('error');
		}
    });
});
