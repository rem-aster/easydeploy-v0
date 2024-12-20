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
					primary: "#02761e",
				},
			},
			{
				dark: {
					...require("daisyui/src/theming/themes")["dark"],
					primary: "#02761e",
				},
			},
		]
	}
}
