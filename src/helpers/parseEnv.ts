export default function parseEnv(contents: string): string
{
	const regex = /\${(.*?)\}/gmi;
	const matches = contents.match(regex);

	if(!process?.env) return contents;
	if(!matches) return contents;

	for(const match of matches)
	{
		const envVar = match.slice(2, -1);

		if(envVar in process.env)
			contents = contents.replace(match, process.env[envVar]);
	}

	return contents;
}
