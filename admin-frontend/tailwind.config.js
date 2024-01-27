import defaultTheme from 'tailwindcss/defaultTheme';
/** @type {import('tailwindcss').Config} */
export default {
	content: [
		'./src/**/*.{html,js,svelte,ts}',
		'./node_modules/flowbite-svelte/**/*.{html,js,svelte,ts}',
	],
	plugins: [require('flowbite/plugin')],
	darkMode: 'class',
	theme: {
		fontFamily: {
			sans: ['Inter', 'Open Sans', ...defaultTheme.fontFamily.sans],
		},
		extend: {},
	},
	plugins: [],
};
