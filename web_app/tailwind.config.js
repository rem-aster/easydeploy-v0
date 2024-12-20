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
					primary: "#2C31FF",
				},
			},
			{
				dark: {
					...require("daisyui/src/theming/themes")["dark"],
					primary: "#2C31FF",
				},
			},
		]
	}
}
