import { Reset } from '@graftini/bricks';

export default function MyApp({ Component, pageProps }) {
  return (
    <>
      <Reset />
      <Component {...pageProps} />
    </>
  );
}
