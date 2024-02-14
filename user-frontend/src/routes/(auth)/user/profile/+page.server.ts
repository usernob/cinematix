import type { ApiResponse } from '@/lib/types/apiResponse';
import type { User } from '@/lib/types/modelTypes';
import { routeApi } from '@/lib/util';
import { fail, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
  default: async ({ request, locals, fetch }) => {
    const formData: FormData = await request.formData();
    const nama = formData.get('nama') as string;
    const email = formData.get('email') as string;
    const avatar = formData.get('avatar') as File;
    const data = new FormData();
    if (avatar?.size && avatar.size > 0) {
      data.append('avatar', avatar);
    }
    data.append('email', email);
    data.append('nama', nama);
    if (!nama || !email) {
      return fail(400, { missing: true, name: nama, email });
    }

    const updateUser = await fetch(routeApi('user/profile/update'), {
      method: 'PUT',
      body: data,
    });

    const res: ApiResponse<{ user?: User } | null> = await updateUser.json();
    if (res.status !== 'ok') {
      return fail(400, { error: true, message: res.message });
    }

    locals.user = res.data?.user;
    return { success: true, user: res.data?.user };
  },
};
