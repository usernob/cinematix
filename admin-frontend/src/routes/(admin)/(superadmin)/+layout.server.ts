import { redirect } from '@sveltejs/kit';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async ({ parent }) => {
	const { user } = await parent();

	if (!user) {
		throw redirect(302, '/login');
	}

	if (user.role !== 'superadmin') {
		throw redirect(302, '/dashboard/petugas');
	}
};
