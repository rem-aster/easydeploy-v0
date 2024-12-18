/** @type {import('tailwindcss').Config} */
module.exports = {
	content: ["internal/views/**/*.templ"],
	theme: {
		extend: {},
	},
	plugins: [require("@tailwindcss/typography"), require("daisyui")],
	daisyui: {
		themes: ["light", "dark"]
	}
}
