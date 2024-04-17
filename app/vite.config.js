import vituum from 'vituum'
import pug from '@vituum/vite-plugin-pug'
import tailwindcss from '@vituum/vite-plugin-tailwindcss'

export default {
    plugins: [vituum(), tailwindcss(), pug({
        root: './src'
    })]
}
