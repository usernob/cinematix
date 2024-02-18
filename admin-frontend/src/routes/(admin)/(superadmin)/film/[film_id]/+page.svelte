<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import Label from 'flowbite-svelte/Label.svelte';
	import Input from 'flowbite-svelte/Input.svelte';
	import Textarea from 'flowbite-svelte/Textarea.svelte';
	import PenSolid from 'flowbite-svelte-icons/PenSolid.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import { invalidateAll, goto } from '$app/navigation';
	import type { ActionData, PageData } from './$types';
	import { routeApi } from '@/lib/util';
	import ActionModal from '@/components/actionModal.svelte';

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
				modalMessage = 'Film Berhasil Diupdate';
			} else {
				modalSuccess = false;
				modalMessage = 'Film Gagal Diupdate';
			}
		}
	}
>
	<div class="flex items-start gap-4">
		<div>
			<Label class="space-y-2">
				<span>Poster</span>
				<div class="relative cursor-pointer overflow-hidden">
					<img
						src={routeApi(data.film.poster_path)}
						alt={data.film.title}
						class="aspect-[9.2/13] h-auto w-52 max-w-max rounded-lg object-cover shadow-lg md:w-64"
					/>
					<div class="group absolute inset-0 flex items-center justify-center hover:bg-black/50">
						<PenSolid class="hidden h-6 w-6 text-white group-hover:block" />
					</div>
					<input type="file" name="poster" class="hidden" />
				</div>
			</Label>
		</div>
		<div class="flex-1">
			<Label class="mb-4 space-y-2">
				<span>Judul Film</span>
				<Input
					type="text"
					placeholder="Something..."
					size="md"
					value={data.film.title}
					name="title"
					required
				/>
			</Label>
			<Label class="mb-4 space-y-2">
				<span>Rating</span>
				<Input
					type="text"
					pattern="\d+(\.\d+)?"
					size="md"
					value={data.film.rating}
					name="rating"
					required
				/>
			</Label>
			<Label class="mb-4 space-y-2">
				<span>Tanggal Rilis</span>
				<Input
					type="date"
					size="md"
					value={new Date(data.film.tanggal_rilis).toISOString().split('T')[0]}
					name="tanggal_rilis"
					required
				/>
			</Label>
		</div>
	</div>
	<Label for="sinopsis" class="mb-2 mt-4">Sinopsis</Label>
	<Textarea
		id="sinopsi"
		placeholder="Pada zaman dahulu..."
		rows="4"
		name="sinopsis"
		value={data.film.sinopsis}
		required
	/>
	{#if form?.message}
		<div class="text-md mb-2 block w-full rounded-lg border-2 border-red-400 bg-red-400/30 p-2.5">
			<p class="text-red-500">{form?.message}</p>
		</div>
	{/if}

	<Button type="submit" class="mt-8">Update</Button>
	<ActionModal
		bind:modalState
		bind:success={modalSuccess}
		bind:message={modalMessage}
		onConfirm={async () => {
			await invalidateAll();
			await goto('/film');
		}}
	/>
</form>
