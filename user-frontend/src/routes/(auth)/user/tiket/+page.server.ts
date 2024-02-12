import type { ApiResponse } from '@/lib/types/apiResponse';
import type { PageServerLoad } from './$types';
import type { Tiket } from '@/lib/types/modelTypes';
import { routeApi } from '@/lib/util';

export const load: PageServerLoad = async ({ parent, fetch }) => {
  const { user } = await parent();
  const getTiket = await fetch(routeApi('user/pesanan'));
  const tiket: ApiResponse<Tiket[]> = await getTiket.json();
  return { user, title: 'Tiket', tiket: tiket.data };
};
