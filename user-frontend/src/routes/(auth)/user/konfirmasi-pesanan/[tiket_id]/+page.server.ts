import type { PageServerLoad } from './$types';
import type { ApiResponse } from '@/lib/types/apiResponse';
import { routeApi } from '@/lib/util';
import type { Film, Kursi, Penayangan, Tiket } from '@/lib/types/modelTypes';

export const load: PageServerLoad = async ({ params, cookies }) => {
  const token = cookies.get('token');

  const res = await fetch(routeApi('user/pesanan/' + params.tiket_id), {
    method: 'GET',
    credentials: 'include',
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  const data: ApiResponse<
    Film & { penayangan: Penayangan & { tiket: Tiket & { kursi: Kursi[] }[] }[] }
  > = await res.json();

  console.log(data.data.penayangan[0].tiket)
  return {
    dataFilm: data.data,
    title: 'Konfirmasi Pesanan',
    listPembayaran: listPembayaran(),
  };
};

const listPembayaran = (): { name: string; src: string }[] => {
  return [
    {
      name: 'dana',
      src: 'https://upload.wikimedia.org/wikipedia/commons/7/72/Logo_dana_blue.svg',
    },
    {
      name: 'gopay',
      src: 'https://upload.wikimedia.org/wikipedia/commons/8/86/Gopay_logo.svg',
    },
    {
      name: 'ovo',
      src: 'https://upload.wikimedia.org/wikipedia/commons/e/eb/Logo_ovo_purple.svg',
    },
  ];
};
