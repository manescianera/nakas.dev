<script>
  let message = 'meanwhile..'
  let href = '/news'
</script>

<a {href}>{message}</a>

<style>
  a {
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);

    opacity: 1;
    text-decoration: none;
    font-size: calc(0.75em + 1vw);
    color: var(--global-color-white);
    border-radius: var(--global-border-radius);
    padding: var(--global-padding);
    animation: appear 2000ms;
  }

  a:hover {
    color: var(--global-color-red);
    background: var(--global-color-bg-secondary);
    text-decoration: underline;
  }

  @keyframes appear {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }
</style>
