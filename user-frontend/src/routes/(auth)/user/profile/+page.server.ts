import type { ApiResponse } from '@/lib/types/apiResponse';
import type { User } from '@/lib/types/modelTypes';
import { routeApi } from '@/lib/util';
import { fail, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
	default: async ({ request, locals, fetch }) => {
		const formData = await request.formData();
		const nama = formData.get('nama');
		const email = formData.get('email');

		console.log({ avatar: formData.get('avatar'), nama, email });
		if (!nama || !email) {
			return fail(400, { missing: true, name: nama, email });
		}

		const updateUser = await fetch(routeApi('user/profile/update'), {
			method: 'POST',
			headers: {
				'Content-Type': 'multipart/form-data',
			},
			body: JSON.stringify({ email, nama, avatar: formData.get('avatar') }),
		});

		const res: ApiResponse<{ user?: User } | null> = await updateUser.json();
    console.log(res)
		if (res.status !== 'ok') {
			return fail(400, { error: true, message: res.message });
		}

		locals.user = res.data?.user;
		return { success: true, user: res.data?.user };
	},
};
