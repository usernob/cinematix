<script>
	import { Carousel } from 'flowbite-svelte';
	let images = [
		{
			alt: 'Cosmic timetraveler',
			src: 'https://source.unsplash.com/1600x900/?beach',
			title: 'cosmic-timetraveler-pYyOZ8q7AII-unsplash.com',
		},
		{
			alt: 'Cristina Gottardi',
			src: 'https://source.unsplash.com/1600x900/?beach',
			title: 'cristina-gottardi-CSpjU6hYo_0-unsplash.com',
		},
		{
			alt: 'Johannes Plenio',
			src: 'https://source.unsplash.com/1600x900/?beach',
			title: 'johannes-plenio-RwHv7LgeC7s-unsplash.com',
		},
		{
			alt: 'Jonatan Pie',
			src: 'https://source.unsplash.com/1600x900/?beach',
			title: 'jonatan-pie-3l3RwQdHRHg-unsplash.com',
		},
		{
			alt: 'Mark Harpur',
			src: 'https://source.unsplash.com/1600x900/?beach',
			title: 'mark-harpur-K2s_YE031CA-unsplash',
		},
		{
			alt: 'Pietro De Grandi',
			src: 'https://source.unsplash.com/1600x900/?beach',
			title: 'pietro-de-grandi-T7K4aEPoGGk-unsplash',
		},
		{
			alt: 'Sergey Pesterev',
			src: 'https://source.unsplash.com/1600x900/?beach',
			title: 'sergey-pesterev-tMvuB9se2uQ-unsplash',
		},
		{
			alt: 'Solo travel goals',
			src: 'https://source.unsplash.com/1600x900/?beach',
			title: 'solotravelgoals-7kLufxYoqWk-unsplash',
		},
	];
</script>

<div class="w-full">
	<Carousel {images} duration="3000" />
</div>
