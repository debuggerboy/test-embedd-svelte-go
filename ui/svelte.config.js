import preprocess from "svelte-preprocess";
import adapter from "@sveltejs/adapter-static";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://svelte.dev/docs/kit/integrations
	// for more information about preprocessors
  kit: {
    adapter: adapter({
      pages: "./../server/uibuild",
      assets: "./../server/uibuild",
      fallback: "index.html",
    }),
  },

  preprocess: [
    preprocess({
      postcss: true,
    }),
  ],
};

export default config;
