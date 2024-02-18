<script lang="ts">
	import Label from 'flowbite-svelte/Label.svelte';
	import Input from 'flowbite-svelte/Input.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import { applyAction, enhance } from '$app/forms';
	import type { ActionData } from './$types';
	import { invalidateAll } from '$app/navigation';
	import ActionModal from '@/components/actionModal.svelte';

	let modalState: boolean = false;
	let modalMessage: string = '';
	let modalSuccess: boolean = false;

	export let form: ActionData;
</script>

<form
	method="POST"
	use:enhance={() =>
		async ({ result }) => {
			await applyAction(result);

			modalState = true;
			if (result.type === 'success') {
				modalSuccess = true;
				modalMessage = 'Petugas baru ditambahkan';
			} else {
				modalSuccess = false;
				modalMessage = 'Petugas gagal ditambahkan';
			}
		}}
>
	<h2 class="mb-8 text-2xl font-bold">Tambah Petugas</h2>
	<div class="mb-6">
		<Label for="nama" class="text-md mb-2 block">Nama</Label>
		<Input
			type="text"
			id="nama"
			placeholder="username"
			class="text-md border-2"
			value={form?.nama ?? ''}
			autocomplete="nama"
			name="nama"
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
			value={form?.email ?? ''}
			autocomplete="email"
			name="email"
			required
		/>
	</div>
	<div class="mb-6">
		<Label for="Password" class="text-md mb-2 block">Password</Label>
		<Input
			id="password"
			placeholder="•••••••••"
			class="text-md border-2"
			value={form?.password ?? ''}
			name="password"
			required
		></Input>
	</div>
	{#if form?.message}
		<div class="text-md mb-2 block w-full rounded-lg border-2 border-red-400 bg-red-400/30 p-2.5">
			<p class="text-red-500">{form?.message}</p>
		</div>
	{/if}
	<ActionModal
		bind:modalState
		bind:success={modalSuccess}
		bind:message={modalMessage}
		onConfirm={async () => await invalidateAll()}
	/>
	<Button type="submit">Buat Akun</Button>
</form>
