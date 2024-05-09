import { defineConfig } from 'unocss';
import presetUno from '@unocss/preset-uno';
import presetIcons from '@unocss/preset-icons';
import presetWebFonts from '@unocss/preset-web-fonts'
import transformerDirectives from '@unocss/transformer-directives';

export default defineConfig({
  theme: {
    colors: {
      primary: '#5463FF',
      'primary-foreground': '#EEEEEE',
      danger: '#FF1818',
      'danger-foreground': '#EEEEEE',
      black: '#222831',
    },
  },
  presets: [
    presetUno(),
    presetIcons(),
    presetWebFonts({
      provider: 'google',
      fonts: {
        sans: 'Poppins:400',
        'sans-bold': 'Poppins:700',
      },
    }),
  ],
  transformers: [
    transformerDirectives(),
  ],
});
