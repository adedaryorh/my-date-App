import 'package:celebut/core/core.dart';
import 'package:flutter/material.dart';

class AppSearchBar extends StatefulWidget {
  const AppSearchBar(
      {super.key, required this.searchCtr, this.onChanged, this.labelText});

  final TextEditingController searchCtr;

  final ValueChanged<String>? onChanged;
  final String? labelText;

  @override
  State<AppSearchBar> createState() => _AppSearchBarState();
}

class _AppSearchBarState extends State<AppSearchBar> {
  @override
  Widget build(BuildContext context) {
    return Container(
      alignment: Alignment.center,
      height: 48,
      width: double.maxFinite,
      decoration: const BoxDecoration(
        color: Color(0xffF8F9FA),
        borderRadius: BorderRadius.all(
          Radius.circular(8),
        ),
      ),
      child: TextFormInput(
        controller: widget.searchCtr,
        decoration: InputDecoration(
          contentPadding: const EdgeInsets.all(10),
          prefixIcon: IconButton(
            onPressed: () {},
            icon: const Icon(
              Icons.search,
              weight: 24,
            ),
          ),
          border: InputBorder.none,
          focusedBorder: InputBorder.none,
          enabledBorder: InputBorder.none,
          hintText: widget.labelText,
          hintStyle: context.textTheme.bodyLarge?.copyWith(
            fontWeight: FontWeight.w600,
            color: const Color(0xff363B4B),
          ),
          fillColor: Colors.transparent,
        ),
        onChanged: widget.onChanged,
      ),
    );
  }
}
