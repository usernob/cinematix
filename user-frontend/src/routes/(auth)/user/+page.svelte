<script lang="ts">
	import { routeApi } from '@/lib/util';
	import Avatar from 'flowbite-svelte/Avatar.svelte';
	import type { PageData } from './$types';

	export let data: PageData;
</script>

<div class="container">
	<div class="pt-4">
		<Avatar
			id="user-menu"
			class="m-auto shadow-lg"
			size="xl"
			border
			src={data.user.avatar ? routeApi(data.user.avatar) : ''}
		/>
		<h3 class="mt-4 text-center text-3xl font-bold">Welcome {data.user.nama}</h3>
	</div>
	<div
		class="m-auto mt-6 flex max-w-[40rem] flex-col items-center justify-start gap-4 text-center text-xl"
	>
		<a href="/user/profile" class="w-full rounded-md border border-gray-500 px-4 py-2">Profil</a>
		<a href="/user/tiket" class="w-full rounded-md border border-gray-500 px-4 py-2">Tiket</a>
		<a
			href="/user/logout"
			class="w-full rounded-md border border-red-500 px-4 py-2 text-red-500 hover:bg-red-500 hover:text-white"
			>Log Out</a
		>
	</div>
</div>
