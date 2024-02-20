<script lang="ts">
	import { applyAction, enhance } from '$app/forms';
	import Label from 'flowbite-svelte/Label.svelte';
	import Input from 'flowbite-svelte/Input.svelte';
	import Textarea from 'flowbite-svelte/Textarea.svelte';
	import PenSolid from 'flowbite-svelte-icons/PenSolid.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import Img from 'flowbite-svelte/Img.svelte';
	import Search from 'flowbite-svelte/Search.svelte';
	import ActionModal from '@/components/actionModal.svelte';
	import CloseCircleSolid from 'flowbite-svelte-icons/CloseCircleSolid.svelte';
	import PlusSolid from 'flowbite-svelte-icons/PlusSolid.svelte';
	import SearchOutline from 'flowbite-svelte-icons/SearchOutline.svelte';
	import { invalidateAll, goto } from '$app/navigation';
	import type { ActionData } from './$types';
	import { routeApi } from '@/lib/util';
	import type { Genre } from '@/lib/types/modelTypes';
	import Modal from 'flowbite-svelte/Modal.svelte';
	import type { ApiResponse } from '@/lib/types/apiResponse';

	export let form: ActionData;

	let genres: Genre[] | undefined;
	let queryGenre: string = '';
	let resGenre: Genre[] | null = null;

	let modalState: boolean = false;
	let modalMessage: string = '';
	let modalSuccess: boolean = false;

	let posterInput: HTMLInputElement;

	let formGenreModal: boolean = false;

	let posterImg: string = '';

	const posterChange = () => {
		if (!posterInput.files) return;
		const file: File = posterInput.files[0];
		const reader = new FileReader();

		reader.readAsDataURL(file);
		reader.onload = (e) => {
			if (e.target) {
				posterImg = e.target.result as string;
			}
		};
	};

	const deleteGenre = (id: number) => {
		genres = genres?.filter((genre) => genre.id !== id);
	};

	const addGenre = (genre: Genre) => {
		if (genres?.some((g) => g.id === genre.id)) {
			return;
		}

		resGenre = resGenre!.filter((item) => item.id !== genre.id);

		if (!genres) {
			genres = [genre];
			return;
		}

		genres = [...genres, genre];
	};

	const findGenre = async (nama: string) => {
		const uppercaseNama: string = nama.toLowerCase();
		const req = await fetch(routeApi('genre/search?' + new URLSearchParams({ q: uppercaseNama })), {
			headers: {
				'Content-Type': 'application/x-www-form-urlencoded',
			},
		});

		const res: ApiResponse<Genre[] | null> = await req.json();
		if (res.status === 'ok' && res.data) {
			resGenre = res.data;
		}
	};
</script>

<form
	method="POST"
	enctype="multipart/form-data"
	use:enhance={() =>
		async ({ result }) => {
			await applyAction(result);

			modalState = true;
			if (result.type === 'success') {
				modalSuccess = true;
				modalMessage = 'Film Berhasil Ditambahkan';
			} else {
				modalSuccess = false;
				modalMessage = 'Film Gagal Ditambahkan';
			}
		}}
>
	<div class="flex flex-col items-start gap-4 md:flex-row">
		<div>
			<Label class="space-y-2">
				<span>Poster</span>
				<div class="relative cursor-pointer overflow-hidden rounded-lg">
					<Img
						src={posterImg}
						figClass="w-full"
						imgClass="aspect-[9.2/13] h-auto w-52 max-w-max object-cover rounded-lg md:w-64"
					/>
					<div class="group absolute inset-0 flex items-center justify-center hover:bg-black/50">
						<PenSolid class="hidden h-6 w-6 text-white group-hover:block" />
					</div>
					<input
						type="file"
						name="poster"
						class="hidden"
						required
						bind:this={posterInput}
						on:change={posterChange}
					/>
				</div>
			</Label>
		</div>
		<div class="w-full flex-1">
			<Label class="mb-4 space-y-2">
				<span>Judul Film</span>
				<Input type="text" placeholder="Something..." size="md" name="title" required />
			</Label>
			<Label class="mb-4 space-y-2">
				<span>Rating</span>
				<Input type="text" pattern="\d+(\.\d+)?" size="md" name="rating" required />
			</Label>
			<Label class="mb-4 space-y-2">
				<span>Tanggal Rilis</span>
				<Input type="date" size="md" name="tanggal_rilis" required />
			</Label>
			<Label class="mb-4 space-y-2">
				<span>Genre</span>
				<div class="justify-start-items-center flex flex-wrap gap-2">
					{#if genres}
						{#each genres as genre (genre.id)}
							<div
								class="flex inline-flex items-center items-center justify-center gap-2 rounded-lg bg-primary-700 px-4 py-2.5 text-center text-sm font-medium text-white hover:bg-primary-800 dark:bg-primary-600 dark:hover:bg-primary-700"
							>
								<input
									type="checkbox"
									class="absolute inset-0 hidden"
									name="genres[]"
									value={genre.id}
									checked
								/>
								<p>{genre.nama}</p>
								<button type="button" on:click={() => deleteGenre(genre.id)}>
									<CloseCircleSolid class="h-5 w-5" />
								</button>
							</div>
						{/each}
					{/if}
					<Button type="button" on:click={() => (formGenreModal = true)}>
						<PlusSolid class="h-5 w-5" />
					</Button>
				</div>
			</Label>
		</div>
	</div>
	<Label for="sinopsis" class="mb-2 mt-4">Sinopsis</Label>
	<Textarea id="sinopsi" placeholder="Pada zaman dahulu..." rows="4" name="sinopsis" required />
	{#if form?.message}
		<div class="text-md mb-2 block w-full rounded-lg border-2 border-red-400 bg-red-400/30 p-2.5">
			<p class="text-red-500">{form?.message}</p>
		</div>
	{/if}

	<Button type="submit" class="mt-8">Tambah</Button>
	<Button
		type="button"
		color="alternative"
		on:click={async () => {
			await invalidateAll();
			await goto('/film');
		}}
		class="mt-8">Batal</Button
	>
	<Modal
		bind:open={formGenreModal}
		backdropClass="fixed inset-0 z-[80] bg-gray-900 bg-opacity-50 dark:bg-opacity-80"
		dialogClass="fixed top-0 start-0 end-0 h-modal md:inset-0 md:h-full z-[90] w-full p-4 flex"
		title="Tambah Genre"
		dismissable
	>
		<Label class="mb-4 space-y-2">
			<span>Masukkan Nama Genre</span>
			<div class="flex gap-2">
				<Search size="md" bind:value={queryGenre} />
				<Button class="!p-2.5" on:click={async () => await findGenre(queryGenre)}>
					<SearchOutline class="h-5 w-5" />
				</Button>
			</div>
		</Label>
		{#if resGenre}
			{#if resGenre.length > 0}
				<div class="flex w-full flex-col divide-y divide-gray-300 dark:divide-gray-500">
					{#each resGenre as genre}
						<div class="flex w-full items-center justify-between py-4">
							<p>{genre.nama}</p>
							<button type="button" on:click={() => addGenre(genre)}>
								<PlusSolid class="h-4 w-4 text-primary-700" />
							</button>
						</div>
					{/each}
				</div>
			{:else}
				<p>Genre tidak ditemukan</p>
			{/if}
		{/if}
	</Modal>

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
