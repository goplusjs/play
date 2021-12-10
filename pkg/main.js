function _main() {
  for (var f in $load) {
    $load[f]()
  }
  $mainPkg = $packages["main"],
    $synthesizeMethods(),
    $packages.runtime.$init(),
    $go($mainPkg.$init, []),
    $flushConsole()
}
_main()
