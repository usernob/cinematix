import { routeApi } from '@/lib/util';
import type { LayoutServerLoadEvent } from './$types';
import type { User } from '@/lib/types/modelTypes';
import type { ApiResponse } from '@/lib/types/apiResponse';
import { redirect } from '@sveltejs/kit';

export async function load(event: LayoutServerLoadEvent) {
  const token = event.cookies.get('token');

  if (!token) {
    return { user: null };
  }
  const getData = await fetch(routeApi('user/info'), {
    headers: { Authorization: `Bearer ${token}` },
  });

  const res: ApiResponse<User> = await getData.json();

  if (res.status === 'ok' || res.data) {
    event.locals.user = res.data;
  } else {
    throw redirect(302, '/login');
  }
  const user = event.locals?.user;
  if (!user) return { user: null };
  return { user };
}
