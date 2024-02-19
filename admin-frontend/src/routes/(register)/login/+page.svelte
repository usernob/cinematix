<script lang="ts">
	import Label from 'flowbite-svelte/Label.svelte';
	import Input from 'flowbite-svelte/Input.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import EyeOutline from 'flowbite-svelte-icons/EyeOutline.svelte';
	import EyeSlashOutline from 'flowbite-svelte-icons/EyeSlashOutline.svelte';
	import type { ActionData } from './$types';
	import { applyAction, enhance } from '$app/forms';
	import { goto, invalidateAll } from '$app/navigation';
	import { session } from '@/lib/stores/session';

	let show: boolean = false;

	export let form: ActionData;
</script>

<svelte:head>
	<title>Login</title>
</svelte:head>
<form
	method="POST"
	use:enhance={() =>
		async ({ result }) => {
			await applyAction(result);

			if (result.type === 'success') {
				const user = result.data?.user;
				if (user) $session.user = user;
				await invalidateAll();
				await goto('/');
			}
		}}
>
	<h2 class="mb-8 text-2xl font-bold">Login</h2>
	<div class="mb-6">
		<Label for="email" class="text-md mb-2 block">Email</Label>
		<Input
			type="email"
			id="email"
			placeholder="john.doe@company.com"
			class="text-md border-2"
			value={form?.email ?? ''}
			autocomplete="email"
			name="email"
			required
		/>
	</div>
	<div class="mb-6">
		<Label for="Password" class="text-md mb-2 block">Password</Label>
		<Input
			type={show ? 'text' : 'password'}
			id="password"
			placeholder="•••••••••"
			class="text-md border-2"
			value={form?.password ?? ''}
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
	{#if form?.incorrect}
		<div class="text-md mb-2 block w-full rounded-lg border-2 border-red-400 bg-red-400/30 p-2.5">
			<p class="text-red-500">Email atau password salah</p>
		</div>
	{/if}
	<Button type="submit">Login</Button>
</form>
