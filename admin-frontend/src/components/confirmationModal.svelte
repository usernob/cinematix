<script lang="ts">
	import Modal from 'flowbite-svelte/Modal.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import QuestionCircleSolid from 'flowbite-svelte-icons/QuestionCircleSolid.svelte';
	export let modalState: boolean = false;
	export let message: string = '';

	export let onConfirm = async () => {};
</script>

<Modal
	bind:open={modalState}
	size="xs"
	dismissable={true}
	backdropClass="fixed inset-0 z-[80] bg-gray-900 bg-opacity-50 dark:bg-opacity-80"
	dialogClass="fixed top-0 start-0 end-0 h-modal md:inset-0 md:h-full z-[90] w-full p-4 flex"
>
	<div class="text-center">
		<QuestionCircleSolid class="mx-auto mb-4 h-20 w-20 text-primary-600" />
		<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">{message}</h3>
		<Button color={'alternative'} on:click={() => (modalState = false)}>Batal</Button>
		<Button
			on:click={async () => {
				modalState = false;
				await onConfirm();
			}}>OK</Button
		>
	</div>
</Modal>
