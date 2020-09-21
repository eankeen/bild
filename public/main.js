export default function main() {
	const portal = document.createElement('portal')
	portal.src = 'https://en.wikipedia.org/wiki/World_Wide_Web'
	document.body.appendChild(portal)

	portal.activate()

	console.info('f')
}
