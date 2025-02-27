import 'package:flutter/material.dart';

class PersistentTabController extends ChangeNotifier {
  PersistentTabController({int initialIndex = 0})
      : _index = initialIndex,
        assert(
          initialIndex >= 0,
          'check that initialIndex is not less than zero',
        );

  int get index => _index;
  int _index;

  set index(int value) {
    assert(value >= 0, 'check that value is not less than zero');
    if (_index == value) {
      return;
    }
    _index = value;
    notifyListeners();
  }
}
