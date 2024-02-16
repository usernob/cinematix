<script lang="ts">
	import Label from 'flowbite-svelte/Label.svelte';
	import Input from 'flowbite-svelte/Input.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import Helper from 'flowbite-svelte/Helper.svelte';
	import EyeOutline from 'flowbite-svelte-icons/EyeOutline.svelte';
	import EyeSlashOutline from 'flowbite-svelte-icons/EyeSlashOutline.svelte';
	import type { ActionData } from '../login/$types';
	import { applyAction, enhance } from '$app/forms';
	import { goto } from '$app/navigation';
	import { session } from '@/lib/stores/session';

	let show: boolean = false;
	let show2: boolean = false;

	let username: string | FormDataEntryValue = '';
	let email: string | FormDataEntryValue = '';
	let password: string | FormDataEntryValue = '';
	let passwordverif: string = '';

	export let form: ActionData;
</script>

<svelte:head>
	<title>Create Account</title>
</svelte:head>
<form
	method="POST"
	use:enhance={() =>
		async ({ result }) => {
			await applyAction(result);

			if (result.type === 'success') {
				const user = result.data?.user;
				if (user) $session.user = user;
				await goto('/user');
			}
		}}
>
	<h2 class="mb-8 text-2xl font-bold">Create Account</h2>
	<div class="mb-6">
		<Label for="username" class="text-md mb-2 block">Username</Label>
		<Input
			type="text"
			id="username"
			placeholder="Your Awesome Username"
			class="text-md border-2"
			name="username"
			autocomplete="username"
			bind:value={username}
			required
		/>
	</div>
	<div class="mb-6">
		<Label for="email" class="text-md mb-2 block">Email</Label>
		<Input
			type="email"
			id="email"
			placeholder="john.doe@company.com"
			class="text-md border-2"
			name="email"
			autocomplete="email"
			bind:value={email}
			required
		/>
	</div>
	<div class="mb-6">
		<Label for="password" class="text-md mb-2 block">Password</Label>
		<Input
			type={show ? 'text' : 'password'}
			id="password"
			placeholder="•••••••••"
			class="text-md border-2"
			bind:value={password}
			name="password"
			required
		>
			<button slot="left" on:click={() => (show = !show)} class="pointer-events-auto">
				{#if !show}
					<EyeOutline class="h-6 w-6" />
				{:else}
					<EyeSlashOutline class="h-6 w-6" />
				{/if}
			</button>
		</Input>
	</div>
	<div class="mb-6">
		<Label
			for="password-verification"
			class="text-md mb-2 block"
			color={password === passwordverif ? 'gray' : 'red'}>Password Verification</Label
		>
		<Input
			type={show2 ? 'text' : 'password'}
			id="password-verification"
			placeholder="•••••••••"
			class="text-md border-2"
			bind:value={passwordverif}
			color={password === passwordverif ? 'base' : 'red'}
			required
		>
			<button slot="left" on:click={() => (show2 = !show2)} class="pointer-events-auto">
				{#if !show2}
					<EyeOutline class="h-6 w-6" />
				{:else}
					<EyeSlashOutline class="h-6 w-6" />
				{/if}
			</button>
		</Input>
		{#if password !== passwordverif}
			<Helper color="red">Password tidak cocok</Helper>
		{/if}
	</div>
	{#if form?.incorrect}
		<div class="text-md mb-2 block w-full rounded-lg border-2 border-red-400 bg-red-400/30 p-2.5">
			<p class="text-red-500">Email sudah terdaftar</p>
		</div>
	{/if}
	<p class="mb-4">
		Sudah punya akun? <a
			href="/login"
			class="text-primary-700 hover:underline dark:text-primary-600">login sekarang</a
		>
	</p>
	<Button type="submit" disabled={password !== passwordverif || password == '' || email == ''}
		>Sign In</Button
	>
</form>
