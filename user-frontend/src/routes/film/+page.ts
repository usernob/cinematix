import type { PageLoad } from './$types';

export const load: PageLoad<{
	images: { alt: string; src: string; title: string }[];
}> = async () => {
	return {
		images: [
			{
				alt: 'Cosmic timetraveler',
				src: 'https://source.unsplash.com/1600x900/?beach',
				title: 'cosmic-timetraveler-pYyOZ8q7AII-unsplash.com',
			},
			{
				alt: 'Cristina Gottardi',
				src: 'https://source.unsplash.com/1600x900/?computer',
				title: 'cristina-gottardi-CSpjU6hYo_0-unsplash.com',
			},
			{
				alt: 'Johannes Plenio',
				src: 'https://source.unsplash.com/1600x900/?people',
				title: 'johannes-plenio-RwHv7LgeC7s-unsplash.com',
			},
			{
				alt: 'Jonatan Pie',
				src: 'https://source.unsplash.com/1600x900/?lanscape',
				title: 'jonatan-pie-3l3RwQdHRHg-unsplash.com',
			},
			{
				alt: 'Mark Harpur',
				src: 'https://source.unsplash.com/1600x900/?monk',
				title: 'mark-harpur-K2s_YE031CA-unsplash',
			},
			{
				alt: 'Pietro De Grandi',
				src: 'https://source.unsplash.com/1600x900/?rain',
				title: 'pietro-de-grandi-T7K4aEPoGGk-unsplash',
			},
			{
				alt: 'Sergey Pesterev',
				src: 'https://source.unsplash.com/1600x900/?sea',
				title: 'sergey-pesterev-tMvuB9se2uQ-unsplash',
			},
			{
				alt: 'Solo travel goals',
				src: 'https://source.unsplash.com/1600x900/?city',
				title: 'solotravelgoals-7kLufxYoqWk-unsplash',
			},
		],
	};
};
