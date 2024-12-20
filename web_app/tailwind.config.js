/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["internal/views/**/*.templ"],
	theme: {
		extend: {},
	},
	plugins: [require("@tailwindcss/typography"), require("daisyui")],
	daisyui: {
		//themes: ["light", "dark"]
		themes: [
			{
				light: {
					...require("daisyui/src/theming/themes")["light"],
					primary: "#c1121f",
				},
			},
			{
				dark: {
					...require("daisyui/src/theming/themes")["dark"],
					primary: "#c1121f",
				},
			},
		]
	}
}
