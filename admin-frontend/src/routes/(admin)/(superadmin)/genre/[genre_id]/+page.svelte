<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import Label from 'flowbite-svelte/Label.svelte';
	import Input from 'flowbite-svelte/Input.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import ActionModal from '@/components/actionModal.svelte';
	import { invalidateAll, goto } from '$app/navigation';
	import type { ActionData, PageData } from './$types';

	export let data: PageData;
	export let form: ActionData;

	let modalState: boolean = false;
	let modalMessage: string = '';
	let modalSuccess: boolean = false;
</script>

<form
	method="POSt"
	enctype="multipart/form-data"
	use:enhance={() =>
		async ({ result }) => {
			await applyAction(result);

			modalState = true;
			if (result.type === 'success') {
				modalSuccess = true;
				modalMessage = 'Genre Berhasil Diupdate';
			} else {
				modalSuccess = false;
				modalMessage = 'Genre Gagal Diupdate';
			}
		}}
>
	<div class="flex flex-col items-start gap-4 md:flex-row">
		<div class="w-full flex-1">
			<Label class="mb-4 space-y-2">
				<span>Nama Genre</span>
				<Input
					type="text"
					placeholder="Something..."
					size="md"
					value={data.genre.nama}
					name="nama"
					required
				/>
			</Label>
		</div>

		<Button type="submit" class="mt-8">Update</Button>
		<Button
			type="button"
			color="alternative"
			on:click={async () => {
				await invalidateAll();
				await goto('/genre');
			}}
			class="mt-8">Batal</Button
		>
		<ActionModal
			bind:modalState
			bind:success={modalSuccess}
			bind:message={modalMessage}
			onConfirm={async () => {
				await invalidateAll();
				await goto('/genre');
			}}
		/>
	</div>
	{#if form?.message}
		<div class="text-md mb-2 block w-full rounded-lg border-2 border-red-400 bg-red-400/30 p-2.5">
			<p class="text-red-500">{form?.message}</p>
		</div>
	{/if}
</form>
