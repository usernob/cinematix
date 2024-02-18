<script lang="ts">
	import Modal from 'flowbite-svelte/Modal.svelte';
	import Button from 'flowbite-svelte/Button.svelte';
	import CheckCircleSolid from 'flowbite-svelte-icons/CheckCircleSolid.svelte';
	import CloserCircleSolid from 'flowbite-svelte-icons/CloseCircleSolid.svelte';
	export let modalState: boolean = false;
	export let success: boolean = false;
	export let message: string = '';

	export let onConfirm = async () => {};
</script>

<Modal bind:open={modalState} size="xs" dismissable={true}>
	<div class="text-center">
		{#if success}
			<CheckCircleSolid class="mx-auto mb-4 h-20 w-20 text-green-400 dark:text-green-200" />
		{:else}
			<CloserCircleSolid class="mx-auto mb-4 h-20 w-20 text-red-400 dark:text-red-200" />
		{/if}
		<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">{message}</h3>
		<Button
			on:click={async () => {
				modalState = false;
				await onConfirm();
			}}>OK</Button
		>
	</div>
</Modal>
