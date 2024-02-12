<script lang="ts">
	import Input from 'flowbite-svelte/Input.svelte';
	import Label from 'flowbite-svelte/Label.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import Avatar from 'flowbite-svelte/Avatar.svelte';
	import PenSolid from 'flowbite-svelte-icons/PenSolid.svelte';
	import { routeApi } from '@/lib/util';
	import type { ActionData, PageData } from './$types';
	import { applyAction, enhance } from '$app/forms';
	import { goto, invalidateAll } from '$app/navigation';
	import { session } from '@/lib/stores/session';

	export let data: PageData;

	export let form: ActionData;

	let input: HTMLInputElement;

	$: avatar = data.user.avatar ? routeApi(data.user.avatar) : '';
	const avatarChange = () => {
		if (!input.files) return;
		const file: File = input.files[0];
		const reader = new FileReader();

		reader.readAsDataURL(file);
		reader.onload = (e) => {
			if (e.target) {
				avatar = e.target.result as string;
			}
		};
	};
</script>

<form
	class="container max-w-[40rem]"
	method="POST"
  enctype="multipart/form-data"
	use:enhance={() =>
		async ({ result }) => {
			await applyAction(result);

			if (result.type === 'success') {
				const user = result.data?.user;
				if (user) $session.user = user;
        await invalidateAll()
				await goto('/user');
			}
		}}
>
	<div class="relative m-auto w-fit">
		<Avatar
			id="user-menu"
			class="object-cover object-center opacity-80 shadow-lg"
			size="xl"
			border
			src={avatar}
		/>
		<div class="absolute inset-0 bottom-0 right-0 flex items-center justify-center">
			<PenSolid class="h-6 w-6 text-white" />
		</div>
		<input
			type="file"
			name="avatar"
			id="avatar"
			class="absolute left-0 top-0 h-full w-full cursor-pointer opacity-0"
			bind:this={input}
			on:change={avatarChange}
		/>
	</div>
	<div class="mb-6">
		<Label for="email" class="mb-2">Username</Label>
		<Input
			type="text"
			name="nama"
			id="nama"
			placeholder="Username"
			required
			value={data.user.nama}
		/>
	</div>
	<div class="mb-6">
		<Label for="email" class="mb-2">Alamat Email</Label>
		<Input
			type="email"
			name="email"
			id="email"
			placeholder="john.doe@company.com"
			required
			value={data.user.email}
		/>
	</div>
	{#if form?.missing}
		<div class="text-md mb-2 block w-full rounded-lg border-2 border-red-400 bg-red-400/30 p-2.5">
			<p class="text-red-500">Harap isi semua data</p>
		</div>
	{/if}

	{#if form?.error}
		<div class="text-md mb-2 block w-full rounded-lg border-2 border-red-400 bg-red-400/30 p-2.5">
			<p class="text-red-500">{form.message}</p>
		</div>
	{/if}
	<Button type="submit">Update</Button>
</form>
