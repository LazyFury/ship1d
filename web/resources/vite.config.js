import { defineConfig } from 'vite';
import babel from '@rollup/plugin-babel'

export default defineConfig({
  plugins: [],
  build: {
    manifest:"manifest.json",
    outDir: '../../public',
    emptyOutDir: true,
    sourcemap: true,
    minify: true,
    target: [
      'es2020',
      'chrome100'
    ],
    modulePreload: {
      polyfill: true
    },
    assetsInlineLimit: 0,
    rollupOptions: {
      input: ['src/main.js'],
      plugins: [
        babel({
          babelHelpers: 'bundled',
          presets: [['@babel/preset-env', {
            "corejs": 3,
            "useBuiltIns": "entry",
            targets: {
              chrome: "100",
              edge: "100",
              firefox: "100",
              safari: "100",
              ie: "11"
            }
          }]],
          exclude: 'node_modules/**',
        })
      ]
    }
  },

});