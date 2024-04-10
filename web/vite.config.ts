import { defineConfig, loadEnv } from "vite";
import vue from "@vitejs/plugin-vue";
import * as path from "path";
import AutoImport from "unplugin-auto-import/vite";
import Components from "unplugin-vue-components/vite";
import { ElementPlusResolver } from "unplugin-vue-components/resolvers";

// https://vitejs.dev/config/
export default defineConfig(({ command, mode }) => {
	const envDir = "./env",
		envPrefix = "VITE_";
	const env = loadEnv(mode, path.resolve(process.cwd(), envDir), envPrefix);
	// console.log("mode", mode);
	// console.log(env.VITE_BASE_URL_API);

	return {
		envDir: envDir,
		envPrefix: envPrefix, // string | string[] | undefined // @default 'VITE_'
		plugins: [
			vue(),
			AutoImport({
				resolvers: [ElementPlusResolver()],
			}),
			Components({
				resolvers: [ElementPlusResolver()],
			}),
		],
		resolve: {
			alias: {
				"~": path.resolve(__dirname, "src"),
			},
			// extensions: ['.js', '.json', '.ts']
		},
		// 基本路径
		// base: '/sdcx'
		server: {
			port: 5000,
			host: true, // http://192.168.1.185:5000/
			proxy: {
				// 本地代理设置，正向代理在前端开发时解决跨域
				"/API": {
					target: env.VITE_BASE_URL,
					changeOrigin: true,
					rewrite: (path) => path.replace(/^\/API/, ""), // 不可以省略rewrite
				},
			},
		},
	};
});
