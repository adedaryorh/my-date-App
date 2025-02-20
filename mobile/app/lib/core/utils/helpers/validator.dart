String? validatePassword(String? value) {
  if (value!.isEmpty) {
    return '';
  } else if (value.length < 8) {
    return 'Must be at least 8 characters';
  }
  return null;
}

String? validateUsername(String? input) {
  if (input == null || input.trim().isEmpty) {
    return 'Username is required';
  } else {
    return null;
  }
}

String? validateName(String? value) {
  const pattern = r"^[.!#$%&'*+<>:;,%@()(/=?^_`{|}~-]";
  final regExp = RegExp(pattern);
  if (value!.isEmpty) {
    return 'Fullname is required';
  }
  if (regExp.hasMatch(value)) {
    return 'invalid';
  }
  return null;
}

String? validateEmailAddress(String? input) {
  const emailRegex =
      r"""^[a-zA-Z0-9.a-zA-Z0-9.!#$%&'*+-/=?^_`{|}~]+@[a-zA-Z0-9]+\.[a-zA-Z]+""";

  if (input == null ||
      input.trim().isEmpty ||
      !RegExp(emailRegex).hasMatch(input)) {
    return 'Please enter a valid email address';
  } else {
    return null;
  }
}

String? validatePhoneNumber(String? value) {
  String number = '${value?.replaceAll(RegExp(r'\s+'), '')}';
  if (value == null) {
    return '';
  } else if (number.startsWith('0') && number.length <= 10) {
    return 'Phone number must be 11 characters';
  } else if (!number.startsWith('0') && number.length <= 9) {
    return 'Phone number must be 10 characters';
  }
  return null;
}
