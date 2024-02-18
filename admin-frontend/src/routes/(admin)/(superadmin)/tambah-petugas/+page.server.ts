import type { ApiResponse } from '@/lib/types/apiResponse';
import type { User } from '@/lib/types/modelTypes';
import { routeApi } from '@/lib/util';
import { fail, type Actions } from '@sveltejs/kit';

export const actions: Actions = {
  default: async ({ request, fetch }) => {
    const formData = await request.formData();
    const nama = formData.get('nama');
    const email = formData.get('email');
    const password = formData.get('password');

    if (!nama || !email || !password) {
      return fail(400, {
        missing: true,
        message: 'All fields are required',
        nama,
        email,
        password,
      });
    }

    const getRegister = await fetch(routeApi('admin/register'), {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        nama,
        email,
        password,
      }),
    });

    const res: ApiResponse<{ token?: string; user?: User } | null> = await getRegister.json();

    console.log(res);

    if (res.status !== 'ok') {
      return fail(400, { error: true, message: res.message });
    }

    return { success: true };
  },
};
