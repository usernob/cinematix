import { dataFilm, seatPesanan, type FilmData } from '@/lib/stores/data';
import type { PageLoad } from './$types';
import { get } from 'svelte/store';
import { redirect } from '@sveltejs/kit';

export const load: PageLoad = async () => {
  if (!get(dataFilm) || !get(seatPesanan)) {
    throw redirect(302, '/');
  }

  return {
    dataFilm: get(dataFilm) as FilmData,
    seatPesanan: get(seatPesanan),
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
      name:"ovo",
      src: 'https://upload.wikimedia.org/wikipedia/commons/e/eb/Logo_ovo_purple.svg'
    }
  ];
};
