import 'package:celebut/apps/personal/profile/presentation/widgets/profile_avatar.dart';
import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:go_router/go_router.dart';
import 'package:intl/intl.dart';

class EditProfile extends StatefulWidget {
  const EditProfile({super.key});

  @override
  State<EditProfile> createState() => _EditProfileState();
}

class _EditProfileState extends State<EditProfile> {
  final TextEditingController fullnameCtr = TextEditingController();
  final TextEditingController bioCtr = TextEditingController();
  final TextEditingController emailCtr = TextEditingController();
  final TextEditingController addressCtr = TextEditingController();
  final TextEditingController phoneNumberCtr = TextEditingController();
  final TextEditingController date = TextEditingController();

  PhoneNumber number = PhoneNumber(isoCode: 'NG');
  String dialCode = '';

  final _formKey = GlobalKey<FormState>();

  @override
  void dispose() {
    emailCtr.dispose();
    fullnameCtr.dispose();
    bioCtr.dispose();
    addressCtr.dispose();
    phoneNumberCtr.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          onPressed: () {
            context.pop();
          },
          icon: const Icon(
            Icons.arrow_back_ios,
            size: 15,
          ),
        ),
      ),
      body: GestureDetector(
        onTap: () {
          FocusScope.of(context).unfocus();
        },
        behavior: HitTestBehavior.opaque,
        child: SafeArea(
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 20),
            child: Form(
              key: _formKey,
              child: ListView(
                children: [
                  const SizedBox(
                    height: 80,
                  ),
                  const ProfileAvatar(),
                  const SizedBox(
                    height: 20,
                  ),
                  Text(
                    'Full Name',
                    style: context.textTheme.bodyMedium
                        ?.copyWith(fontWeight: FontWeight.w600),
                  ),
                  const SizedBox(
                    height: 5,
                  ),
                  TextFormInput(
                    autoValidateMode: AutovalidateMode.onUserInteraction,
                    controller: fullnameCtr,
                    validator: validateName,
                    labelText: 'Enter full name',
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  Text(
                    'Bio',
                    style: context.textTheme.bodyMedium
                        ?.copyWith(fontWeight: FontWeight.w600),
                  ),
                  const SizedBox(
                    height: 5,
                  ),
                  Container(
                    height: 126,
                    width: double.maxFinite,
                    padding: const EdgeInsets.symmetric(horizontal: 10),
                    decoration: BoxDecoration(
                      border: Border.all(),
                      borderRadius: BorderRadius.circular(10),
                    ),
                    child: TextFormField(
                      controller: bioCtr,
                      maxLines: 4,
                      decoration: InputDecoration(
                        enabledBorder: InputBorder.none,
                        border: InputBorder.none,
                        focusedBorder: InputBorder.none,
                        errorBorder: InputBorder.none,
                        disabledBorder: InputBorder.none,
                        filled: false,
                        hintText: 'Enter bio',
                        hintStyle: context.textTheme.bodyMedium?.copyWith(
                          fontWeight: FontWeight.w600,
                          color: Colors.black45,
                        ),
                      ),
                    ),
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  Text(
                    'Date of Birth',
                    style: context.textTheme.bodyMedium
                        ?.copyWith(fontWeight: FontWeight.w600),
                  ),
                  const SizedBox(
                    height: 5,
                  ),
                  TextFormInput(
                    readOnly: true,
                    controller: date,
                    validator: validateDate,
                    labelText: 'DD/MM/YYYY',
                    suffixIcon: const Icon(
                      Icons.calendar_month,
                      color: Colors.black,
                      size: 20,
                    ),
                    inputFormatters: [
                      FilteringTextInputFormatter.deny(RegExp('[ ]')),
                    ],
                    onTap: () async {
                      final value = await showPlatformDatePicker(
                        context: context,
                        initialDate: DateTime.now(),
                        firstDate: DateTime.now()
                            .subtract(const Duration(days: 365 * 50)),
                        lastDate: DateTime.now(),
                      );
                      if (value == null) return;
                      final formatter = DateFormat('yyyy-MM-dd');
                      final formattedDate = formatter.format(value);
                      setState(() {
                        date.text = formattedDate;
                      });
                    },
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  Text(
                    'Phone Number',
                    style: context.textTheme.bodyMedium
                        ?.copyWith(fontWeight: FontWeight.w600),
                  ),
                  const SizedBox(
                    height: 5,
                  ),
                  InternationalPhoneNumberInput(
                    onInputChanged: (PhoneNumber number) {
                      dialCode = '${number.dialCode}';
                    },
                    autoValidateMode: AutovalidateMode.onUserInteraction,
                    onInputValidated: (bool value) {},
                    initialValue: number,
                    textFieldController: phoneNumberCtr,
                    keyboardType: const TextInputType.numberWithOptions(
                      signed: true,
                      decimal: true,
                    ),
                    validator: validatePhoneNumber,
                    onPressed: () {},
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  Text(
                    'Email',
                    style: context.textTheme.bodyMedium
                        ?.copyWith(fontWeight: FontWeight.w600),
                  ),
                  const SizedBox(
                    height: 5,
                  ),
                  TextFormInput(
                    autoValidateMode: AutovalidateMode.onUserInteraction,
                    controller: emailCtr,
                    validator: validateEmailAddress,
                    labelText: 'Enter email',
                  ),
                  const SizedBox(
                    height: 10,
                  ),
                  Text(
                    'Address',
                    style: context.textTheme.bodyMedium
                        ?.copyWith(fontWeight: FontWeight.w600),
                  ),
                  const SizedBox(
                    height: 5,
                  ),
                  TextFormInput(
                    autoValidateMode: AutovalidateMode.onUserInteraction,
                    controller: addressCtr,
                    validator: validateName,
                    labelText: 'Enter address',
                  ),
                  const SizedBox(
                    height: 30,
                  ),
                ],
              ),
            ),
          ),
        ),
      ),
    );
  }
}
