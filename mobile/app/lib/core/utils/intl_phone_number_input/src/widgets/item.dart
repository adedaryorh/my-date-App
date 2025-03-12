import 'package:celebut/core/utils/intl_phone_number_input/src/models/country_model.dart';
import 'package:celebut/core/utils/intl_phone_number_input/src/utils/util.dart';
import 'package:flutter/material.dart';

/// [Item]
class Item extends StatelessWidget {
  final Country? country;
  final bool? showFlag;
  final bool? useEmoji;
  final TextStyle? textStyle;
  final bool withCountryNames;
  final double? leadingPadding;
  final bool trailingSpace;

  const Item({
    Key? key,
    this.country,
    this.showFlag,
    this.useEmoji,
    this.textStyle,
    this.withCountryNames = false,
    this.leadingPadding = 10,
    this.trailingSpace = true,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    String dialCode = (country?.dialCode ?? '');
    if (trailingSpace) {
      dialCode = dialCode.padRight(2, '   ');
    }
    return SizedBox(
      child: Row(
        mainAxisAlignment: MainAxisAlignment.start,
        mainAxisSize: MainAxisSize.min,
        children: <Widget>[
          // SizedBox(width: leadingPadding),
          // _Flag(
          //   country: country,
          //   showFlag: showFlag,
          //   useEmoji: useEmoji,
          // ),
          const SizedBox(width: 10.0),
          Text(
            dialCode,
            textDirection: TextDirection.ltr,
            style: Theme.of(context).textTheme.bodyMedium,
          ),
          const SizedBox(width: 5.0),
          //Todo ---> Remove drop down icon
          const Icon(
            Icons.keyboard_arrow_down,
            color: Colors.black,
            size: 16,
          ),
          const SizedBox(width: 10.0),
        ],
      ),
    );
  }
}

// ignore: unused_element
class _Flag extends StatelessWidget {
  final Country? country;
  final bool? showFlag;
  final bool? useEmoji;

  // ignore: unused_element
  const _Flag({Key? key, this.country, this.showFlag, this.useEmoji})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return country != null && showFlag!
        ? Container(
            child: useEmoji!
                ? Text(
                    Utils.generateFlagEmojiUnicode(country?.alpha2Code ?? ''),
                    style: Theme.of(context)
                        .textTheme
                        .bodyMedium
                        ?.copyWith(fontSize: 17),
                  )
                : Image.asset(
                    country!.flagUri,
                    width: 32.0,
                    errorBuilder: (context, error, stackTrace) {
                      return const SizedBox.shrink();
                    },
                  ),
          )
        : const SizedBox.shrink();
  }
}
