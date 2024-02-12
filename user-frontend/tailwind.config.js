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
		container: {
			center: true,
			padding: '1rem',
		},
		fontFamily: {
			sans: ['Inter', 'Open Sans', ...defaultTheme.fontFamily.sans],
		},
		extend: {
			colors: {
				// flowbite-svelte
				primary: {
					50: '#FFF5F2',
					100: '#FFF1EE',
					200: '#FFE4DE',
					300: '#FFD5CC',
					400: '#FFBCAD',
					500: '#FE795D',
					600: '#EF562F',
					700: '#EB4F27',
					800: '#CC4522',
					900: '#A5371B',
				},
			},
		},
	},
};
