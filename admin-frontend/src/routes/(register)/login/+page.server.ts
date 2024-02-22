import type { ApiResponse } from '@/lib/types/apiResponse';
import type { User } from '@/lib/types/modelTypes';
import { routeApi } from '@/lib/util';
import { fail, type Actions, redirect } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load: PageServerLoad = async (event) => {
  const user = event.locals.user;
  if (user) {
    if (user.role === 'admin') {
      throw redirect(302, '/dashboard/petugas');
    } else {
      throw redirect(302, '/dashboard/superadmin');
    }
  }
};

export const actions = {
  default: async ({ request, cookies, locals }) => {
    const formData = await request.formData();
    const email = formData.get('email');
    const password = formData.get('password');

    if (!email || !password) {
      return fail(400, { missing: true, email, password });
    }

    const getLogin = await fetch(routeApi('login/admin'), {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email,
        password,
      }),
    });
    const res: ApiResponse<{ token?: string; user?: User } | null> = await getLogin.json();

    if (res.status !== 'ok' || !res.data?.token || !res.data?.user) {
      return fail(400, { incorrect: true, email, password });
    }

    cookies.set('admin-token', res.data?.token, {
      path: '/',
      maxAge: 60 * 60 * 24, // one day
    });

    delete res.data?.user.password;

    locals.user = res.data.user;
    return { success: true, user: res.data?.user };
  },
} satisfies Actions;
